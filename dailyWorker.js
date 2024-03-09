import { Client } from 'pg';

// Replace these values with your PostgreSQL connection details
const connectionString = "postgresql://yakshit:-eZfWw2zQKffFmvntDaL-g@sparkdb-6147.6xw.aws-ap-southeast-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full";

const client = new Client({
    connectionString: connectionString,
});

async function resetColumns() {
    try {
        await client.connect();
        // Execute the update query to reset 'hits' and 'status' columns in 'UserDetails' table
        const query1 = ` 
            UPDATE userrequests SET hits = 0;
            UPDATE userkeystatus SET status = 'ok';
        `;
        const result = await client.query(query1);
        console.log('Columns reset successfully:', result.rowCount);
    } catch (error) {
        console.error('Error resetting columns:', error);
    } finally {
        await client.end();
    }
}

resetColumns();
