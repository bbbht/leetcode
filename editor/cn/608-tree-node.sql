package main
// 608 tree-node 2022-08-01 11:27:07

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select id,
    case when p_id is null then 'Root'
         when id in (select p_id from tree) then 'Inner'
         else 'Leaf'
     end as `type`
from tree

#leetcode submit region end(Prohibit modification and deletion)
