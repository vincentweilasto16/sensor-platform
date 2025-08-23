-- name: InsertSensorData :exec
INSERT INTO sensor_data (sensor_type, sensor_value, device_code, device_number, timestamp)
VALUES (?, ?, ?, ?, ?);

-- name: GetSensorDataByDeviceCodeAndNumber :many
SELECT *
FROM sensor_data
WHERE device_code = ?
    AND device_number = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountSensorDataByDeviceCodeAndNumber :one
SELECT COUNT(*)
FROM sensor_data
WHERE device_code = ? AND device_number = ?;;

-- name: GetSensorDataByTime :many
SELECT *
FROM sensor_data
WHERE timestamp BETWEEN ? AND ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetSensorDataByDeviceAndTime :many
SELECT *
FROM sensor_data
WHERE device_code = ?
    AND device_number = ?
    AND timestamp BETWEEN ? AND ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateSensorData :exec
UPDATE sensor_data
SET sensor_value = ?
WHERE device_code = ?
    AND device_number = ?
    AND timestamp = ?;
    
-- name: DeleteSensorData :exec
DELETE FROM sensor_data
WHERE device_code = ?
    AND device_number = ?
    AND timestamp = ?;