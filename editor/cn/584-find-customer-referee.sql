package main
// 584 find-customer-referee 2022-08-01 11:22:43

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select name
from customer
where referee_id <> 2 or referee_id is null;

#leetcode submit region end(Prohibit modification and deletion)
