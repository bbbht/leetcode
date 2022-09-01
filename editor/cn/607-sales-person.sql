package main
// 607 sales-person 2022-08-01 11:25:36

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select name
from SalesPerson
where sales_id not in(
	select o.sales_id
	from Orders as o
	left join Company as c
	on o.com_id=c.com_id
	where c.name='RED'
);

#leetcode submit region end(Prohibit modification and deletion)
