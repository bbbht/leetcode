# 换座位 2022-07-29 17:45:38

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT
	s1.id,
	ifnull( s2.student, s1.student ) AS student
FROM
	seat AS s1
	LEFT JOIN ( SELECT IF ( id % 2 = 0, id - 1, id + 1 ) AS id, student FROM seat ) AS s2 ON s1.id = s2.id
ORDER BY
	id ASC

#leetcode submit region end(Prohibit modification and deletion)
