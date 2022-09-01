package main
// 586 customer-placing-the-largest-number-of-orders 2022-08-01 11:23:48

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select
    customer_number
from orders
group by customer_number
order by count(customer_number) desc
limit 1

#leetcode submit region end(Prohibit modification and deletion)
