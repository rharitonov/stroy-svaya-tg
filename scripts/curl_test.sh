curl -X POST http://localhost:8080/insertpdrline \
  -H "Content-Type: application/json" \
  -d '{
    "project_id": 1,
	"pile_field_id": 1,
    "pile_number": "399",
    "start_date": "2025-04-05T14:30:00Z",
    "fact_pile_head": 10.75,
	"recorded_by": "Иванов Иван Иванович"
  }'
  
  
 curl http://localhost:8080/insertpdrlines?project_id=1