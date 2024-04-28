CREATE TABLE IF NOT EXISTS taxation_rules (
	id SERIAL PRIMARY KEY,
	rule_name VARCHAR(255),
	min_income DECIMAL(10, 2),
	max_income DECIMAL(10, 2),
	tax_rate INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tax_configurations (
	id SERIAL PRIMARY KEY,
	condition_name VARCHAR(30) NOT NULL,
	amount DECIMAL(10, 2) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO taxation_rules (id,rule_name,min_income, max_income, tax_rate) VALUES
(1,'0-150,000',0,150000,0),
(2,'150,001-500,000',150000,500000,10),
(3,'500,001-1,000,000',500000,1000000,15),
(4,'1,000,001-2,000,000',1000000,2000000,20),
(5,'2,000,001 ขึ้นไป',2000000,'infinity'::numeric,35);

INSERT INTO tax_configurations (condition_name, amount) VALUES 
('donation', 100000),
('personal', 60000),
('k-receipt', 50000);