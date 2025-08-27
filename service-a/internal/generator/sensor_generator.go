package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"service-a/internal/constants"
	"service-a/internal/domain"
	"service-a/internal/messaging"
)

type Generator struct {
	mu        sync.Mutex
	frequency time.Duration
	ticker    *time.Ticker
	broker    messaging.Broker
	stop      chan struct{}
	update    chan time.Duration
	running   bool
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewGenerator(freq time.Duration, broker messaging.Broker) *Generator {
	return &Generator{
		frequency: freq,
		broker:    broker,
		stop:      make(chan struct{}),
		update:    make(chan time.Duration),
	}
}

// Start the generator if not already running
func (g *Generator) Start() {
	g.mu.Lock()
	if g.running {
		g.mu.Unlock()
		return
	}
	g.running = true
	g.ticker = time.NewTicker(g.frequency)
	g.mu.Unlock()

	go func() {
		for {
			select {
			case <-g.ticker.C:
				data := g.generateSensorData()
				if err := g.publishSensor(data); err != nil {
					fmt.Println("❌ failed to publish:", err)
				} else {
					fmt.Println("✅ published sensor data:", data)
				}
			case newFreq := <-g.update:
				g.mu.Lock()
				g.ticker.Stop()
				g.frequency = newFreq
				g.ticker = time.NewTicker(newFreq)
				g.mu.Unlock()
				fmt.Println("⏱ frequency updated to:", newFreq)
			case <-g.stop:
				g.mu.Lock()
				if g.ticker != nil {
					g.ticker.Stop()
				}
				g.running = false
				g.mu.Unlock()
				fmt.Println("🛑 generator stopped")
				return
			}
		}
	}()
}

// Stop the generator
func (g *Generator) Stop() {
	select {
	case g.stop <- struct{}{}:
	default:
		// already stopping/stopped
	}
}

// UpdateFrequency safely updates ticker frequency
func (g *Generator) UpdateFrequency(newFreq time.Duration) {
	select {
	case g.update <- newFreq:
	default:
		// if generator not running, just update freq
		g.mu.Lock()
		g.frequency = newFreq
		if g.ticker != nil {
			g.ticker.Stop()
			g.ticker = time.NewTicker(newFreq)
		}
		g.mu.Unlock()
		fmt.Println("⏱ frequency updated to:", newFreq)
	}
}

// generateSensorData remains the same
func (g *Generator) generateSensorData() domain.SensorData {
	sType := constants.SensorTypes[rnd.Intn(len(constants.SensorTypes))]
	return domain.SensorData{
		SensorType:   sType,
		SensorValue:  setRandomSensorValueByType(sType),
		DeviceCode:   setSensorRandomDeviceCode(1),
		DeviceNumber: int32(rnd.Intn(1000) + 1),
		Timestamp:    time.Now().UTC(),
	}
}

func (g *Generator) publishSensor(data domain.SensorData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return g.broker.Publish(
		context.Background(),
		[]byte(fmt.Sprintf("sensor-%d", rnd.Intn(1000))),
		jsonData,
	)
}
