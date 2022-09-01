package main
// 185 department-top-three-salaries 2022-08-01 11:11:23

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select
	d.name as Department, e.name as Employee, e.salary as Salary
from
	department d, employee e
where e.departmentId = d.id and (
    	select count(distinct salary) From employee where departmentId=d.id and salary > e.salary
	) < 3

#leetcode submit region end(Prohibit modification and deletion)
