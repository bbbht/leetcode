# 有趣的电影 2022-07-29 17:45:21

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT
	*
FROM
	cinema
WHERE
	description != 'boring'
	AND id % 2 != 0
ORDER BY
	rating DESC

#leetcode submit region end(Prohibit modification and deletion)
