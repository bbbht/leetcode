package main
// 184 department-highest-salary 2022-08-01 11:03:49

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select
	d.name as Department,
	e.name as Employee,
	e.salary as Salary
from
	department d,
	employee e
where e.departmentId=d.id and e.salary=(select max(salary) from employee where departmentId=d.id)

#leetcode submit region end(Prohibit modification and deletion)
