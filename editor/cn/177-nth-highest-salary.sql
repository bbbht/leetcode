package main
// 177 nth-highest-salary 2022-08-01 10:56:53

#leetcode submit region begin(Prohibit modification and deletion)
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  Declare M INT;
  Set M = N-1;
  RETURN (
      # Write your MySQL query statement below.
      SELECT
      	( SELECT salary FROM `employee` GROUP BY salary ORDER BY salary DESC LIMIT M, 1 ) AS SecondHighestSalary
  );
END
#leetcode submit region end(Prohibit modification and deletion)
