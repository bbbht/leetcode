#leetcode submit region begin(Prohibit modification and deletion)
# Please write a DELETE statement and DO NOT write a SELECT statement.
# Write your MySQL query statement below

DELETE p1
FROM
	Person AS p1
	LEFT JOIN ( SELECT min( Id ) AS Id, Email FROM Person GROUP BY Email ) AS p2 ON p1.Id = P2.Id
WHERE
	p2.Id IS NULL

#leetcode submit region end(Prohibit modification and deletion)
