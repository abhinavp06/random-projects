import { Pool } from 'pg';
import { faker } from '@faker-js/faker';

const pool = new Pool({
  user: '<USERNAME>',
  host: 'localhost',
  database: 'campaign_gateway',
  password: '<PASSWORD>',
  port: 5432,
});

function generateUserData() {
  return {
    name: faker.person.fullName(),
    age: faker.number.int({ min: 18, max: 80 }),
    email: faker.internet.email(),
    mobile: 'xxxx',
    organization: faker.company.name(),
    joining_date: faker.date.past({ years: 5 }).toISOString().split('T')[0],
    created_at: new Date().toISOString(),
  };
}

async function insertUserData() {
  const client = await pool.connect();
  try {
    await client.query('BEGIN');

    const insertQuery = `
      INSERT INTO public.user_data (name, age, email, mobile, organization, joining_date, created_at)
      VALUES ($1, $2, $3, $4, $5, $6, $7)
      ON CONFLICT (email) DO NOTHING
    `;

    let successfulInserts = 0;
    for (let i = 0; i < 10000; i++) {
      const user = generateUserData();
      try {
        await client.query(insertQuery, [
          user.name,
          user.age,
          user.email,
          user.mobile,
          user.organization,
          user.joining_date,
          user.created_at,
        ]);
        successfulInserts++;
        if (i % 100 === 0) {
          console.log(`Processed ${i} rows, ${successfulInserts} inserted`);
        }
      } catch (error) {
        if (error.code === '23505') {
          console.warn(`Skipping row ${i + 1} due to duplicate email: ${user.email}`);
          continue;
        }
        throw error;
      }
    }

    await client.query('COMMIT');
    console.log(`Successfully inserted ${successfulInserts} rows into user_data table`);

    const result = await client.query('SELECT COUNT(*) FROM public.user_data');
    console.log(`Total rows in user_data table: ${result.rows[0].count}`);
  } catch (error) {
    await client.query('ROLLBACK');
    console.error('Error during data insertion:', {
      message: error.message,
      code: error.code,
      detail: error.detail,
      hint: error.hint,
    });
  } finally {
    client.release();
  }
}

async function main() {
  try {
    const client = await pool.connect();
    try {
      const res = await client.query(`
        SELECT column_name, data_type 
        FROM information_schema.columns 
        WHERE table_schema = 'public' AND table_name = 'user_data'
      `);
      if (res.rows.length === 0) {
        throw new Error('Table public.user_data does not exist');
      }
      console.log('Table schema:', res.rows);
    } finally {
      client.release();
    }
    await insertUserData();
  } catch (error) {
    console.error('Failed to execute script:', {
      message: error.message,
      code: error.code,
    });
  } finally {
    await pool.end();
    console.log('Database connection closed');
  }
}

main();