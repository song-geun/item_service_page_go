
> DB : mysql8.0.41

# **소스코드**

[itemServicePage Client](https://github.com/song-geun/itemServicePage_Client)
[itemServicePage GO 코드](https://github.com/song-geun/item\_service\_page\_go)
[itemServicePage Spring 코드](https://github.com/song-geun/itemServicePage_Server)


#### **t_product 테이블**

![](https://velog.velcdn.com/images/rikoriko/post/936bedfc-7e3e-4cee-95ee-6d49af2a9dc2/image.png)

#### **t_product_data 테이블**

![](https://velog.velcdn.com/images/rikoriko/post/b54b4e42-2970-48c6-a0df-90c2d0573613/image.png)




## **1\. USP\_ProductDataManage**

```
create
    definer = root@localhost procedure USP_ProductDataManage(IN prstype varchar(50), IN Pprod_data_id bigint,
                                                             IN Pp_id bigint, IN Pp_name varchar(50), IN P_value bigint,
                                                             IN P_quantity bigint, IN P_DATE varchar(50))
BEGIN
    IF prstype = 'S1' THEN
        SELECT * FROM t_product_data;
    ELSEIF prstype = 'S2' THEN
        SELECT * FROM t_product_data WHERE DATE = P_DATE;
    ELSEIF prstype = 'S3' THEN
        SELECT * FROM t_product_data WHERE LEFT(DATE,6) = LEFT(P_DATE,6);
    ELSEIF prstype = 'I1' THEN
        IF EXISTS(SELECT 1 FROM t_product_data) THEN
            IF NOT EXISTS(SELECT 1 FROM t_product_data WHERE Pprod_data_id = prod_data_id) THEN
                SET @M1 = (SELECT MAX(p_id) FROM t_product) + 1;
                INSERT INTO t_product_data(prod_data_id, p_id, p_name, value, p_quantity, date) values (@M1,Pp_id, Pp_name, P_value, P_quantity,P_DATE);
            ELSE
                UPDATE t_product_data SET p_id = Pp_id, p_name = Pp_name, value = P_value, p_quantity = P_quantity, date = P_DATE WHERE prod_data_id = Pprod_data_id;
            END IF;
        ELSE
            INSERT INTO t_product_data(prod_data_id, p_id, p_name, value, p_quantity, date) values (1,Pp_id, Pp_name, P_value, P_quantity,P_DATE);
        END if;
    END IF;
END;
```
​
## **2\. USP\_ProductDataManage2**
​
```
create
    definer = root@localhost procedure USP_ProductDataManage2(IN prstype varchar(10), IN Pprod_data_id varchar(20),
                                                              IN Pp_id varchar(5), IN Pp_name varchar(250),
                                                              IN Pvalue varchar(20), IN Pp_quantity varchar(20),
                                                              IN Pp_DATE1 varchar(250), IN Pp_DATE2 varchar(250))
BEGIN
    IF prstype = 'S1' THEN
       SELECT p_id as prod_data_id, p_id,p_name, value, SUM(quantity) as quantity, p_id as date
       FROM t_product_data
       WHERE DATE >= Pp_DATE1 AND DATE <= Pp_DATE2
       GROUP BY p_id,p_name, value;
    END IF;
END;
​
```
​
## **3\. USP\_ProductManage**
​
```
create
    definer = root@localhost procedure USP_ProductManage(IN prstype varchar(50), IN Pp_id bigint,
                                                         IN Pp_name varchar(50), IN P_value bigint,
                                                         IN P_quantity bigint)
BEGIN
    IF prstype = 'S1' THEN
        SELECT * FROM t_product;
    ELSEIF prstype = 'I1' THEN
        IF EXISTS(SELECT * FROM t_product) THEN
            IF NOT EXISTS(SELECT * FROM t_product WHERE p_id = Pp_id) THEN
                SET @M1 = (SELECT MAX(p_id) FROM t_product) + 1;
                INSERT INTO t_product(p_id, p_name, value, quantity) values (@M1,Pp_name, P_value, P_quantity);
            ELSE
                UPDATE t_product SET p_id = Pp_id, p_name = Pp_name, value = P_value, quantity = P_quantity WHERE p_id = Pp_id;
            END IF;
        ELSE
            INSERT INTO t_product(p_id, p_name, value, quantity) values (1, Pp_name, P_value, P_quantity);
        END if;

    ELSEIF prstype = 'D1' THEN
            DELETE FROM t_product WHERE p_id = Pp_id;
    ELSEIF prstype = 'U1' THEN
            UPDATE t_product SET p_id = Pp_id, p_name = Pp_name, value = P_value, quantity = P_quantity;
    END IF;
END;


```
