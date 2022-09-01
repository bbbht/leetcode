package main
// 1050 actors-and-directors-who-cooperated-at-least-three-times 2022-08-01 11:29:18

#leetcode submit region begin(Prohibit modification and deletion)
# Write your MySQL query statement below

select
    a.actor_id,A.director_id
from actorDirector as a
group by a.actor_id , a.director_id
having count(a.timestamp) >=3;

#leetcode submit region end(Prohibit modification and deletion)
