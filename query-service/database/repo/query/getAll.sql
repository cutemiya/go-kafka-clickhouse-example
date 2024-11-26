select sum(price) as price, tripId as tripId from offer
group by tripId;