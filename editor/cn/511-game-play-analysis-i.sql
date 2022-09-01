package main
// 游戏玩法分析 I 2022-07-29 17:47:02

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select
    player_id,
    min(event_date) as 'first_login'
from activity
group by player_id

#leetcode submit region end(Prohibit modification and deletion)
