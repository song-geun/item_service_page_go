DB : mysql8.0.41

[  
](https://github.com/song-geun/item_service_page_go)[https://github.com/song-geun/item\_service\_page\_go](https://github.com/song-geun/item_service_page_go)

[GitHub - song-geun/item\_service\_page\_go

Contribute to song-geun/item\_service\_page\_go development by creating an account on GitHub.

github.com](https://github.com/song-geun/item_service_page_go)

[https://github.com/song-geun/itemServicePage\_Server](https://github.com/song-geun/itemServicePage_Server)

[GitHub - song-geun/itemServicePage\_Server: 엄마용

엄마용. Contribute to song-geun/itemServicePage\_Server development by creating an account on GitHub.

github.com](https://github.com/song-geun/itemServicePage_Server)

DB구조 t\_product

[##_Image|kage@Cf8Xn/btsMXdqBtkF/msgOw8IVvsaLgy8coNcZK1/img.png|CDM|1.3|{"originWidth":1193,"originHeight":49,"style":"alignCenter"}_##]

DB구조 t\_product\_data

[##_Image|kage@dqdS0r/btsMXqcfZpR/NaIW8nLyEZRW87f7hqZYek/img.png|CDM|1.3|{"originWidth":1884,"originHeight":58,"style":"alignCenter"}_##]

1\. USP\_ProductDataManage

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
        SELECT * FROM t_product_data WHERE DATE = left(P_DATE,6);
    ELSEIF prstype = 'I1' THEN
        UPDATE t_product_data SET prod_data_id=Pprod_data_id, p_id =  Pp_id, p_name = Pp_name,
                                  value = P_value, quantity =  P_quantity, DATE = P_DATE;
    END IF;
END;

```

2\. USP\_ProductDataManage2

```
create
    definer = root@localhost procedure USP_ProductDataManage2(IN prstype varchar(10), IN Pprod_data_id varchar(20),
                                                              IN Pp_id varchar(5), IN Pp_name varchar(250),
                                                              IN Pvalue varchar(20), IN Pp_quantity varchar(20),
                                                              IN Pp_DATE1 varchar(250), IN Pp_DATE2 varchar(250))
BEGIN
    IF prstype = 'S1' THEN
       SELECT p_id as prod_data_id, p_id,p_name, value, SUM(p_quantity) as p_quantity, p_id as date
       FROM t_product_data
       WHERE DATE >= Pp_DATE1 AND DATE <= Pp_DATE2
       GROUP BY p_id,p_name, value;
    END IF;
END;

```

3\. USP\_ProductManage

```
create
    definer = root@localhost procedure USP_ProductManage(IN prstype varchar(50), IN Pp_id bigint,
                                                         IN Pp_name varchar(50), IN P_value bigint,
                                                         IN P_quantity bigint)
BEGIN
    IF prstype = 'S1' THEN
        SELECT * FROM t_product;
    ELSEIF prstype = 'I1' THEN
        IF NOT EXISTS(SELECT 1 FROM t_product WHERE p_id = Pp_id LIMIT 1) THEN
            INSERT INTO t_product values (Pp_id, Pp_name, P_value, P_quantity);
        ELSE
            UPDATE t_product SET p_id = Pp_id, p_name = Pp_name, value = P_value, quantity = P_quantity WHERE p_id = Pp_id;
        END IF;
    ELSEIF prstype = 'D1' THEN
        IF NOT EXISTS(SELECT 1 FROM t_product WHERE p_id = Pp_id) THEN
            DELETE FROM t_product WHERE p_id = Pp_id;
        END IF;
    ELSEIF prstype = 'U1' THEN
        IF NOT EXISTS(SELECT 1 FROM t_product WHERE p_id = Pp_id) THEN
            UPDATE t_product SET p_id = Pp_id, p_name = Pp_name, value = P_value, quantity = P_quantity;
        END IF;
    END IF;
END;

```