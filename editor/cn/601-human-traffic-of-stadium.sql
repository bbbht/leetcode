# 体育馆的人流量 2022-07-29 17:44:57

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT DISTINCT
	s3.*
FROM
	stadium AS s3,
	(
SELECT
	id
FROM
	(
SELECT
	@i :=
IF
	( s.`people` >= 100, @i + 1, 0 ),
IF
	( @i > 2, s.`people`, NULL ) AS `people`,
	s.id AS id
FROM
	`stadium` AS s,
	( SELECT @i := 0 ) AS tmp
ORDER BY
	s.id ASC
	) AS s2
WHERE
	people IS NOT NULL
	) AS s4
WHERE
	s4.id = s3.id + 1
	OR s4.id = s3.id + 2
	OR s4.id = s3.id

#leetcode submit region end(Prohibit modification and deletion)
