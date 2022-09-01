package main
// 596 classes-more-than-5-students 2022-08-01 11:24:37

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT class from courses group by class having count(DISTINCT student) >= 5;

#leetcode submit region end(Prohibit modification and deletion)
