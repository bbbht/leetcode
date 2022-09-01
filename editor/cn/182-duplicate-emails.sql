#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT `Email` FROM `Person` GROUP BY `Email` HAVING count(*) > 1

#leetcode submit region end(Prohibit modification and deletion)
