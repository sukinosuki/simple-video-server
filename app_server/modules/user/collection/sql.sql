SELECT vc.vid,
       vc.created_at AS collection_time,
       v.title,
       v.cover
FROM video_collection vc
         LEFT JOIN video v
                   ON vc.vid = v.id
WHERE vc.uid = ?
ORDER BY vc.created_at DESC
LIMIT ?,?