SELECT 
    w2.id
FROM 
    Weather w1, Weather w2
WHERE 
    DATEDIFF(w2.recordDate, w1.recordDate) = 1 
AND 
    w2.temperature > w1.temperature;


-- SELECT 
--     w1.id
-- FROM 
--     Weather w1
-- JOIN 
--     Weather w2
-- ON 
--     DATEDIFF(w1.recordDate, w2.recordDate) = 1
-- WHERE 
--     w1.temperature > w2.temperature;


-- SELECT 
--     w1.id
-- FROM 
--     Weather w1
-- WHERE 
--     w1.temperature > (
--         SELECT 
--             w2.temperature
--         FROM 
--             Weather w2
--         WHERE 
--             w2.recordDate = DATE_SUB(w1.recordDate, INTERVAL 1 DAY)
--     );

-- WITH PreviousWeatherData AS
-- (
--     SELECT 
--         id,
--         recordDate,
--         temperature, 
--         LAG(temperature, 1) OVER (ORDER BY recordDate) AS PreviousTemperature,
--         LAG(recordDate, 1) OVER (ORDER BY recordDate) AS PreviousRecordDate
--     FROM 
--         Weather
-- )
-- SELECT 
--     id 
-- FROM 
--     PreviousWeatherData
-- WHERE 
--     temperature > PreviousTemperature
-- AND 
--     recordDate = DATE_ADD(PreviousRecordDate, INTERVAL 1 DAY);