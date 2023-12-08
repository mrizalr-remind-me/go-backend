CREATE TABLE IF NOT EXISTS todo (                                
	id UUID PRIMARY KEY,                                          
	title VARCHAR(255) NOT NULL,                                  
	description VARCHAR(255),                                     
	remind_at TIMESTAMP DEFAULT NULL,                             
	created_at TIMESTAMP DEFAULT NOW(),                           
	updated_at TIMESTAMP DEFAULT NULL                             
);                                                               
                                                                 
CREATE INDEX IF NOT EXISTS idx_todo_remind_at ON todo(remind_at);