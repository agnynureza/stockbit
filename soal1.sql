select u.id, u.username, s.username as ParentUserName
	from users as u 
	left join users as s on s.id = u.parent
	order by 1;