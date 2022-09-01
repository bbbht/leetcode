#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below
SELECT DISTINCT	`Num` AS `ConsecutiveNums`
FROM
	(SELECT
			@i :=IF( @s = l.`Num`, @i + 1, 0 ),
			@s := l.`Num`,
			IF( @i >= 2, l.`Num`, NULL ) AS `Num`
		FROM
			`Logs` AS l, ( SELECT @i := 0, @s := NULL ) AS tmp
		ORDER BY l.id ASC
	) AS l2
WHERE
	`Num` IS NOT NULL
#leetcode submit region end(Prohibit modification and deletion)
