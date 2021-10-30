create or replace function product_stock(item_id uuid) 
	returns table (productid uuid, stock_akhir int, stock_cart int)
	language plpgsql as
	$$
declare
	ref_cursor refcursor;
	ref_cursor2 refcursor;
	total_cart int default 0;
	total_payment int default 0;
	total_receive int default 0;
	cart_id uuid;
	transaction_stat varchar default '';
	item_qty int default 0;
begin
	stock_akhir := 0;
	stock_cart := 0;
	productid := item_id;
    
	open ref_cursor for select shopping_cart_id, qty from shopping_cart_details
		where cast(to_char(deleted_at,'YYYY') as INT) < 1000 and product_id = item_id order by created_at;
	loop
		fetch ref_cursor into cart_id, item_qty;
		
		exit when not found;
		-- raise notice '% - %', cart_id, item_qty;
		
	open ref_cursor2 for 
		select t2.transaction_status from checkouts t1 
		left join checkout_payments t2 
			on cast(t2.check_out_id as varchar) = cast(t1."id" as varchar) and t2.transaction_status in ('pending', 'captured', 'settlement')
		where t1.shopping_cart_id = cart_id and cast(to_char(t1.deleted_at,'YYYY') as INT) < 1000 order by t1.created_at desc limit 1;
		
	fetch ref_cursor2 into transaction_stat;
	if found then 
		total_payment := total_payment + item_qty;
	else 
		stock_cart := stock_cart + item_qty;
	end if;
	close ref_cursor2;
	end loop;
	close ref_cursor;
	
	open ref_cursor for select sum(qty) from purchase_receiving_details 
		where cast(product_id as varchar) = cast(item_id as varchar) and deleted_by = '';
	fetch ref_cursor into total_receive;
	if total_receive is null then
		total_receive := 0;
	end if;
	close ref_cursor;
		
	stock_akhir := total_receive - total_payment - stock_cart;
		
	-- raise notice '%', stock_akhir;
	return query select productid, stock_akhir, stock_cart;
end;
$$
