-- name: InsertSensorData :exec
INSERT INTO sensor_data (sensor_value, sensor_type, id1, id2, timestamp)
VALUES (?, ?, ?, ?, ?);

-- name: GetSensorDataByID :many
SELECT * FROM sensor_data
WHERE id1 = ? AND id2 = ?;

-- name: GetSensorDataByTime :many
SELECT * FROM sensor_data
WHERE timestamp BETWEEN ? AND ?;

-- name: UpdateSensorData :exec
UPDATE sensor_data
SET sensor_value = ?
WHERE id1 = ? AND id2 = ? AND timestamp = ?;

-- name: DeleteSensorData :exec
DELETE FROM sensor_data
WHERE id1 = ? AND id2 = ? AND timestamp = ?;