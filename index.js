const pg = require("pg")

const connectionString = "db_url";

const client = new pg.Client({
    connectionString: connectionString,
});

exports.handler = async (event) => {
    try {
        await client.connect();
        // Execute the update query to reset 'hits' and 'status' columns in 'UserDetails' table
        const query1 = ` 
            UPDATE userdetails SET hits = 0;
            UPDATE userkeystatus SET status = 'ok';
        `;
        const result = await client.query(query1);
        console.log('Columns reset successfully:', result.rowCount);
    } catch (error) {
        console.error('Error resetting columns:', error);
    } finally {
        await client.end();
    }
};
