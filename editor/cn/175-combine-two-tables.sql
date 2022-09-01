#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below
select p.FirstName, p.LastName, a.City, a.State from Person as p
left join Address as a
on p.PersonId = a.PersonId
#leetcode submit region end(Prohibit modification and deletion)
