-- name: InsertSensorData :exec
INSERT INTO sensor_data (sensor_type, sensor_value, device_code, device_number, timestamp)
VALUES (?, ?, ?, ?, ?);

-- name: GetSensors :many
SELECT *
FROM sensor_data
WHERE 
    (sqlc.narg(device_code) IS NULL OR device_code = sqlc.narg(device_code))
    AND (sqlc.narg(device_number) IS NULL OR device_number = sqlc.narg(device_number))
    AND (sqlc.narg(start_time) IS NULL OR timestamp >= sqlc.narg(start_time))
    AND (sqlc.narg(end_time) IS NULL OR timestamp <= sqlc.narg(end_time))
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountSensors :one
SELECT COUNT(*)
FROM sensor_data
WHERE 
    (sqlc.narg(device_code) IS NULL OR device_code = sqlc.narg(device_code))
    AND (sqlc.narg(device_number) IS NULL OR device_number = sqlc.narg(device_number))
    AND (sqlc.narg(start_time) IS NULL OR timestamp >= sqlc.narg(start_time))
    AND (sqlc.narg(end_time) IS NULL OR timestamp <= sqlc.narg(end_time));
    
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