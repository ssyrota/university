```
psql initialized
{ steps: 107, i: 0 }
Uploading chunk 1. Size: 94.69986057281494 MB. Total: 94.69986057281494 MB.
ingesting 300000 rows into psql
ingesting 300000 rows into clickhouse
ingested 300000 rows into clickhouse in 9435 ms
ingested 300000 rows into psql in 18714 ms
{ steps: 107, i: 1 }
...
{ steps: 107, i: 42 }
Uploading chunk 43. Size: 92.98324680328369 MB. Total: 4017.7345685958862 MB.
ingesting 300000 rows into psql
ingesting 300000 rows into clickhouse
ingested 300000 rows into clickhouse in 9244 ms
```

Let's stop here. We have ingested 12_054_000 rows. The total size of the json data is 4 GB.

The psql volume size is 3.7Gb.

The clickhouse volume size is 563.7 MB.

1. Count the overall quantity of sold products(58051058).

- Psql - 1s
- Clickhouse - 0.074s

2. Count the overall price of sold products.

2.1 Get unique products.

Clickhouse(0.686s) - `SELECT DISTINCT ON (product_id) * FROM sells;`

Psql(3s) - `SELECT DISTINCT ON (product_id) * FROM sells;`

2.2 Backfill prices.

Clickhouse:

```sql
CREATE TABLE prices (price Float64, product_id UUID, PRIMARY KEY product_id) engine=MergeTree() AS SELECT randNormal(10, 2) as price, product_id FROM (SELECT DISTINCT product_id FROM sells);
```

Psql:

```sql
CREATE TABLE prices AS select distinct product_id, floor(random() * 10 + 1)::int FROM public.sells;
```

2.3 Count the overall price of sold products.

Psql(7s):

```sql
explain analyze SELECT
    SUM(p.floor * s.quantity) AS total_price
FROM sells s
INNER JOIN prices p ON s.product_id = p.product_id;
```

Clickhouse(0.147s):

```sql
SELECT sum(price * quantity) AS total_price
FROM sells
INNER JOIN prices USING (product_id)
```

3. Count the overall price of sold products for some period.

Psql(1s):

```sql
explain analyze SELECT
    SUM(p.floor * s.quantity) AS total_price
FROM sells s
INNER JOIN prices p ON s.product_id = p.product_id where s.time BETWEEN '2023-02-01' AND '2023-03-01';
```

Clickhouse(0.09s):

```sql
SELECT sum(price * quantity) AS total_price
FROM sells
INNER JOIN prices USING (product_id)
```

4. Count the overall price of sold products for some period of product in store.

Psql(180ms) index on all props:

```sql
explain analyze SELECT
    SUM(p.floor * s.quantity) AS total_price
FROM sells s
INNER JOIN prices p ON s.product_id = p.product_id where s.time >= '2023-02-01' AND s.time<= '2023-03-01' AND s.product_id='47298028-d0da-4946-a007-08016d55b47d' AND s.store_id='f67f94cd-6f71-48a4-9b32-35122211577d';
```

Clickhouse(100ms) - no index on all props:

```sql
SELECT sum(price * quantity) AS total_price
FROM sells
INNER JOIN prices USING (product_id)
WHERE ((sells.time >= '2023-02-01') AND (sells.time <= '2023-03-01')) AND (sells.product_id = '47298028-d0da-4946-a007-08016d55b47d') AND (sells.store_id = 'f67f94cd-6f71-48a4-9b32-35122211577d')
```

5.Count the overall price of sold products for some period of product in all stores.

Psql(224ms) index on all props:

```sql
explain analyze SELECT
    SUM(p.floor * s.quantity) AS total_price
FROM sells s
INNER JOIN prices p ON s.product_id = p.product_id where s.time >= '2023-02-01' AND s.time<= '2023-03-01' AND s.product_id='47298028-d0da-4946-a007-08016d55b47d';
```

Clickhouse(77ms) - no index on all props:

```sql
SELECT sum(price * quantity) AS total_price
FROM sells
INNER JOIN prices USING (product_id)
WHERE ((sells.time >= '2023-02-01') AND (sells.time <= '2023-03-01')) AND (sells.product_id = '47298028-d0da-4946-a007-08016d55b47d')
```

6.Count the overall price of sold products for some period in all stores(group by store).

Psql(1.4s) index on all props:

```sql
explain analyze SELECT
    store_id,SUM(p.floor * s.quantity) AS total_price
FROM sells s
INNER JOIN prices p ON s.product_id = p.product_id where s.time >= '2023-02-01' AND s.time<= '2023-03-01' GROUP BY store_id;
```

Clickhouse(100ms) - no index on all props:

```sql
SELECT sum(price * quantity) AS total_price
FROM sells
INNER JOIN prices USING (product_id)
WHERE ((sells.time >= '2023-02-01') AND (sells.time <= '2023-03-01')) AND (sells.product_id = '47298028-d0da-4946-a007-08016d55b47d')
```
