select sum(price) as price, tripId as tripId from offer where tripId = ?
group by tripId;