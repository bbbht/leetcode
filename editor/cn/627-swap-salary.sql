# 变更性别 2022-07-29 17:46:00

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select id, name, if(sex = 'm','f','m') sex,salary from salary;

#leetcode submit region end(Prohibit modification and deletion)
