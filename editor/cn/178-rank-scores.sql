#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT
	s.Score,
	CONVERT(t2.Rank, SIGNED) AS Rank
FROM
	Scores AS s
	LEFT JOIN (
SELECT
	@i := @i + 1 AS Rank,
	t1.Score
FROM
	( SELECT Score FROM Scores GROUP BY Score ORDER BY Score DESC ) AS t1,
	( SELECT @i := 0 ) t
	) AS t2 ON s.Score = t2.Score
ORDER BY
	s.Score DESC

#leetcode submit region end(Prohibit modification and deletion)
