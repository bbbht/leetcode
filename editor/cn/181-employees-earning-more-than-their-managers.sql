#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select e1.Name as Employee from Employee as e1 ,Employee as e2 where e1.ManagerId=e2.Id AND e1.Salary>e2.Salary

#leetcode submit region end(Prohibit modification and deletion)
