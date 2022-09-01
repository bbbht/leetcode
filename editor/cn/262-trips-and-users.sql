#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT
	t.`Request_at` AS `Day`,
	CAST(
	( 1- count( IF ( t.`Status` = 'completed', 1, NULL ) ) / count( * ) ) AS DECIMAL ( 3, 2 )
	) AS `Cancellation Rate`
FROM
	`Trips` AS t
WHERE
	t.Client_Id NOT IN ( SELECT `Users_id` FROM `Users` WHERE `Banned` = 'Yes' )
	AND t.Driver_Id NOT IN ( SELECT `Users_id` FROM `Users` WHERE `Banned` = 'Yes' )
	AND t.`Request_at` BETWEEN '2013-10-01'
	AND '2013-10-03'
GROUP BY
	t.`Request_at`

#leetcode submit region end(Prohibit modification and deletion)
