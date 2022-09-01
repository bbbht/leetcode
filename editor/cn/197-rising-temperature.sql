#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

SELECT
	w1.Id AS Id
FROM
	Weather AS w1,
	Weather AS w2
WHERE
    w1.RecordDate = DATE_ADD( w2.RecordDate, INTERVAL 1 DAY )
AND
	w1.Temperature > w2.Temperature

#leetcode submit region end(Prohibit modification and deletion)
