#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select Name as Customers from Customers where Id not in (select CustomerId from Orders)

#leetcode submit region end(Prohibit modification and deletion)
