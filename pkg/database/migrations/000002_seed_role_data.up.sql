INSERT INTO roles (id, name, permissions) VALUES 
(1, 'ADMIN', '{
    "user_management": true,
    "role_management": true,
    "attendance_management": true,
    "leave_management": true,
    "work_schedule_management": true,
    "report_access": true,
    "market_management": true
}'::jsonb),

(2, 'MANAGER', '{
    "user_management": false,
    "role_management": false,
    "attendance_management": true,
    "leave_management": true,
    "work_schedule_management": true,
    "report_access": true,
    "market_management": true
}'::jsonb),

(3, 'EMPLOYEE', '{
    "user_management": false,
    "role_management": false,
    "attendance_management": false,
    "leave_management": false,
    "work_schedule_management": false,
    "report_access": false,
    "market_management": false,
    "self_attendance": true,
    "self_leave": true,
    "self_schedule": true
}'::jsonb);