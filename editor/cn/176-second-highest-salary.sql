package main
// 176 second-highest-salary 2022-08-01 10:46:17

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT (select salary FROM `employee` group by salary order by salary desc limit 1,1) as SecondHighestSalary;

#leetcode submit region end(Prohibit modification and deletion)
