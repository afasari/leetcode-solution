# Write your MySQL query statement below
select c.name
from Customer c
where referee_id is null
or referee_id != 2