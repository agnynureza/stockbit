SELECT u.id, u.username, s.username AS ParentUserName
	FROM users as u 
  LEFT JOIN users as s on s.id = u.parent
  ORDER BY 1;