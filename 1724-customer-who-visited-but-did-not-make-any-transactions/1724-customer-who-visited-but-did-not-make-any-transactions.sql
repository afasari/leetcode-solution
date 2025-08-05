-- # Write your MySQL query statement below

-- select v.customer_id, n.count_no_trans as count_no_trans
-- from Visits v
-- left join (
--     select v1.customer_id, count(v1.visit_id) as count_no_trans, t1.amount 
--     from Visits v1
--     left join Transactions t1
--     on v1.visit_id = t1.visit_id
--     group by v1.customer_id, v1.visit_id
--     having t1.amount is null
-- ) n
-- on n.customer_id = v.customer_id
-- group by v.customer_id, n.customer_id
-- ;

SELECT 
  customer_id, 
  COUNT(visit_id) AS count_no_trans 
FROM 
  Visits 
WHERE 
  visit_id NOT IN (
    SELECT 
      visit_id 
    FROM 
      Transactions
  ) 
GROUP BY 
  customer_id