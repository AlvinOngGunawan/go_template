package constanta

const (
	QueryExportInvoice = "SELECT invoices.id, invoices.no_faktur, invoices.no_ref, invoices.customer_name, " +
		"invoices.sold_at, u.fullname, invoices.status, w.name as warehouse_name, uo.name as office_name, invoices.price_segment, " +
		"id.item_count, invoices.subtotal, invoices.discount, invoices.discount_value, " +
		"invoices.total, id.total_return, invoices.outstanding_payment, invoices.last_payment, invoices.note," +
		"d.name as division_name " +
		"FROM invoices " +
		"JOIN warehouses w ON w.id = invoices.from_warehouse_id " +
		"JOIN users u ON u.id = invoices.confirmed_by " +
		"JOIN offices uo ON uo.id = w.office_id " +
		"JOIN (SELECT " +
		"invoice_id, " +
		"COUNT(*) AS item_count, " +
		"SUM(CASE WHEN invoice_return_id IS NOT NULL THEN total ELSE 0 END) AS total_return " +
		"FROM invoice_details " +
		"GROUP BY invoice_id) id ON id.invoice_id = invoices.id " +
		"JOIN divisions d ON d.id = invoices.division_id " +
		"WHERE invoices.division_id != ? " +
		"AND invoices.status IN (?) " +
		"AND w.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"AND invoices.division_id IN (SELECT division_id FROM users_divisions WHERE user_id = ?) " +
		"AND invoices.sold_at BETWEEN ? AND ? "

	QueryExportInvoiceDetail = "SELECT i.no_faktur, i.no_ref, i.customer_name, i.sold_at, " +
		"id.id, inv.no_faktur_pgi, inv.imei_sn, COALESCE(inv.pawned_at, '-') as pawned_at, inv.created_at, " +
		"inv.source, w.name as warehouse_name, o.name as office_name, ik.name as kind_name, ib.name as brand_name, " +
		"it.name as type_name, inv.year, isp.name as spec_name, (CASE WHEN inv.is_batangan = 1 THEN 'Batangan' ELSE 'Lengkap' END) as batangan, " +
		"inv.grade, inv.grade_pgi, inv.price_at_pawn, id.grade_a_price, inv.adj_grade, inv.adj_spec, inv.adj_batangan, inv.adj_other, " +
		"id.adjustment, id.adjustment_price_segment, id.discount_value, id.total, inv.capital, " +
		"(CASE WHEN id.invoice_return_id IS NULL THEN id.total - inv.capital ELSE 0 END) as pl, " +
		"note.notes as notes, inv.adj_other_note, id.invoice_return_id " +
		"FROM invoice_details id " +
		"JOIN invoices i ON id.invoice_id = i.id " +
		"JOIN inventories inv ON id.inventory_id = inv.id " +
		"JOIN item_types it ON inv.type_id = it.id " +
		"JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN item_kinds ik ON ib.item_kind_id = ik.id " +
		"JOIN warehouses w ON inv.last_warehouse_id = w.id " +
		"JOIN offices o ON w.office_id = o.id " +
		"JOIN item_specs isp ON inv.item_spec_id = isp.id " +
		"LEFT JOIN (SELECT dtl.id, CASE WHEN dtl.note IS NULL OR dtl.note = '' THEN GROUP_CONCAT(adj.name SEPARATOR ', ') " +
		"ELSE dtl.note " +
		"END AS notes " +
		"FROM invoice_details dtl " +
		"JOIN invoice_detail_adjustments adj ON adj.invoice_detail_id = dtl.id " +
		"GROUP BY dtl.id) note ON id.id = note.id " +
		"WHERE i.id IN (?)"

	QueryExportInvoiceDetailAdjustment = "SELECT i.no_faktur, i.customer_name, inv.no_faktur_pgi, ida.name, ida.adjustment " +
		"FROM invoices i " +
		"JOIN invoice_details id ON i.id = id.invoice_id " +
		"JOIN invoice_detail_adjustments ida ON id.id = ida.invoice_detail_id " +
		"JOIN inventories inv ON id.inventory_id = inv.id " +
		"WHERE ida.invoice_detail_id IN (?) " +
		"ORDER BY id.id"

	QueryExportInvoiceReturn = "SELECT ir.no, i.no_faktur, i.customer_name, i.sold_at, inv.no_faktur_pgi, " +
		"i.no_ref, inv.imei_sn,COALESCE(inv.pawned_at, '-') as pawned_at, inv.created_at, " +
		"COALESCE(w.name, '-') as warehouse_name, COALESCE(o.name, '-') as office_name, " +
		"ik.name as kind_name, ib.name as brand_name, it.name as type_name, inv.year, " +
		"COALESCE(isp.name, '-') as spec_name, (CASE WHEN inv.is_batangan = 1 THEN 'Batangan' ELSE 'Lengkap' END) as batangan, " +
		"inv.grade, inv.grade_pgi, id.total, ir.created_at as created_at_return, u.fullname, id.return_reason " +
		"FROM invoice_details id " +
		"JOIN invoices i ON id.invoice_id = i.id " +
		"JOIN inventories inv ON id.inventory_id = inv.id " +
		"JOIN invoice_returns ir ON ir.id = id.invoice_return_id " +
		"JOIN users u ON u.id = ir.created_by " +
		"JOIN item_types it ON inv.type_id = it.id " +
		"JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN item_kinds ik ON ib.item_kind_id = ik.id " +
		"JOIN warehouses w ON inv.last_warehouse_id = w.id " +
		"JOIN offices o ON w.office_id = o.id " +
		"JOIN item_specs isp ON inv.item_spec_id = isp.id " +
		"WHERE id.invoice_return_id IN (?)"

	QueryExportInvoiceGold = "SELECT invoices.id, invoices.no_faktur, invoices.no_ref, invoices.customer_name, " +
		"invoices.sold_at, u.fullname, invoices.status, w.name as warehouse_name, uo.name as office_name, invoices.price_segment, " +
		"id.item_count, invoices.subtotal, invoices.discount, invoices.discount_value, " +
		"invoices.total, id.total_return, invoices.outstanding_payment, invoices.last_payment, invoices.note," +
		"d.name as division_name, invoices.invoice_type " +
		"FROM invoices " +
		"JOIN warehouses w ON w.id = invoices.from_warehouse_id " +
		"JOIN users u ON u.id = invoices.confirmed_by " +
		"JOIN offices uo ON uo.id = w.office_id " +
		"JOIN (SELECT " +
		"invoice_id, " +
		"COUNT(*) AS item_count, " +
		"SUM(CASE WHEN invoice_return_id IS NOT NULL THEN total ELSE 0 END) AS total_return " +
		"FROM invoice_details " +
		"GROUP BY invoice_id) id ON id.invoice_id = invoices.id " +
		"JOIN divisions d ON d.id = invoices.division_id " +
		"WHERE invoices.division_id = ? " +
		"AND invoices.status IN (?) " +
		"AND w.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"AND invoices.division_id IN (SELECT division_id FROM users_divisions WHERE user_id = ?) " +
		"AND invoices.sold_at BETWEEN ? AND ? "

	QueryExportInvoiceGoldDetail = "SELECT i.no_faktur, i.no_ref, i.customer_name, i.sold_at, " +
		"inv.no_faktur_pgi, inv.imei_sn, COALESCE(inv.pawned_at, '-') as pawned_at, inv.created_at, " +
		"inv.source, w.name as warehouse_name, o.name as office_name, ik.name as kind_name, ib.name as brand_name, " +
		"it.name as type_name, ig.purity, ig.dry_weight, ig.weight_reduction, " +
		"(CASE WHEN ig.weight_reduction IS NULL THEN '-' ELSE CAST(ig.dry_weight - ig.weight_reduction AS CHAR) END) AS net_weight, " +
		"gmm.name as gold_mint_mark_name, gt.name, ig.piece_count, id.discount_value, id.total, inv.capital, " +
		"(CASE WHEN id.invoice_return_id IS NULL THEN id.total - inv.capital ELSE 0 END) as pl, " +
		"id.invoice_return_id, inv.item_kind_id, inv.type_id " +
		"FROM invoice_details id " +
		"JOIN invoices i ON id.invoice_id = i.id " +
		"JOIN inventories inv ON id.inventory_id = inv.id " +
		"JOIN item_kinds ik ON inv.item_kind_id = ik.id " +
		"LEFT JOIN item_types it ON inv.type_id = it.id " +
		"LEFT JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN warehouses w ON inv.last_warehouse_id = w.id " +
		"JOIN offices o ON w.office_id = o.id " +
		"LEFT JOIN inventory_golds ig ON ig.inventory_id = inv.id " +
		"LEFT JOIN gold_mint_marks gmm ON ig.gold_mint_mark_id = gmm.id " +
		"LEFT JOIN gold_types gt ON ig.gold_type_id = gt.id " +
		"WHERE i.id IN (?)"

	QueryExportReturnInvoiceGold = "SELECT ir.no, i.no_faktur, i.customer_name, i.sold_at, inv.no_faktur_pgi, " +
		"i.no_ref ,COALESCE(inv.pawned_at, '-') as pawned_at, inv.created_at, " +
		"COALESCE(w.name, '-') as warehouse_name, COALESCE(o.name, '-') as office_name, " +
		"ik.name as kind_name, ib.name as brand_name, it.name as type_name, ig.purity, ig.dry_weight, ig.weight_reduction, " +
		"(CASE WHEN ig.weight_reduction IS NULL THEN '-' ELSE CAST(ig.dry_weight - ig.weight_reduction AS CHAR) END) AS net_weight, " +
		"gmm.name as gold_min_mark_name, gt.name as type_name, ig.piece_count, id.total, ir.created_at as created_at_return, u.fullname, id.return_reason " +
		"FROM invoice_details id " +
		"JOIN invoices i ON id.invoice_id = i.id " +
		"JOIN inventories inv ON id.inventory_id = inv.id " +
		"JOIN invoice_returns ir ON ir.id = id.invoice_return_id " +
		"JOIN users u ON u.id = ir.created_by " +
		"JOIN item_kinds ik ON inv.item_kind_id = ik.id " +
		"LEFT JOIN item_types it ON inv.type_id = it.id " +
		"LEFT JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN warehouses w ON inv.last_warehouse_id = w.id " +
		"JOIN offices o ON w.office_id = o.id " +
		"JOIN item_specs isp ON inv.item_spec_id = isp.id " +
		"LEFT JOIN inventory_golds ig ON ig.inventory_id = inv.id " +
		"LEFT JOIN gold_mint_marks gmm ON ig.gold_mint_mark_id = gmm.id " +
		"LEFT JOIN gold_types gt ON ig.gold_type_id = gt.id " +
		"WHERE id.invoice_return_id IN (?)"

	QueryExportReturnInvoiceReport = "SELECT ir.id, ir.no, i.no_faktur, i.no_ref, i.customer_name, ir.created_at, u.fullname, ir.amount, ir.refund " +
		"FROM invoice_returns ir " +
		"JOIN invoices i ON ir.invoice_id = i.id " +
		"JOIN warehouses w ON w.id = i.from_warehouse_id " +
		"JOIN users u ON u.id = ir.created_by " +
		"WHERE w.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"AND i.division_id IN (SELECT division_id FROM users_divisions WHERE user_id = ?) " +
		"AND ir.created_at BETWEEN CONCAT(?, ' 00:00:00') AND CONCAT(?, ' 23:59:59')"

	QueryExportCustomerTransactionReport = "SELECT cs.created_at, (CASE WHEN ca.city_id IS NULL THEN c.name ELSE CONCAT(c.name, ' ' ,ci.name) END) AS customer_name, " +
		"cs.type, cs.amount, cs.credit_change, cs.credit_after, cs.debt_change, cs.debt_after, u.fullname, cs.params " +
		"FROM customer_transactions cs " +
		"JOIN customers c ON cs.customer_id = c.id " +
		"JOIN users u ON u.id = cs.created_by " +
		"LEFT JOIN customer_addresses ca ON c.id = ca.customer_id " +
		"LEFT JOIN cities ci ON ca.city_id = ci.id " +
		"WHERE cs.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"AND cs.created_at BETWEEN CONCAT(?, ' 00:00:00') AND CONCAT(?, ' 23:59:59')"

	QueryExportDeliveryBatchEXTReport = "SELECT d.id, d.date, db.branch, d.source, inv.item_count " +
		"FROM delivery_batches db " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN (SELECT i.delivery_batch_id , COUNT(*) as item_count " +
		"FROM inventories i " +
		"JOIN divisions_item_kinds dik ON i.item_kind_id = dik.item_kind_id " +
		"WHERE dik.division_id != ? " +
		"AND dik.division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?) " +
		"GROUP BY i.delivery_batch_id) AS inv ON inv.delivery_batch_id = db.id " +
		"WHERE d.source = 'EXT' " +
		"AND d.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"ORDER BY BRANCH"

	QueryExportDeliveryBatchReport = "SELECT d.id, d.date, db.branch, d.source, COALESCE(inv.item_count, 0) AS item_count " +
		"FROM delivery_batches db " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN (SELECT i.delivery_batch_id , COUNT(*) as item_count " +
		"FROM inventories i " +
		"JOIN divisions_item_kinds dik ON i.item_kind_id = dik.item_kind_id " +
		"WHERE dik.division_id != ? " +
		"AND dik.division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?) " +
		"GROUP BY i.delivery_batch_id) AS inv ON inv.delivery_batch_id = db.id " +
		"WHERE d.source != 'EXT' " +
		"AND d.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"AND d.date BETWEEN ? AND ? " +
		"ORDER BY BRANCH"

	QueryExportDeliveryItemEXTReport = "SELECT d.date, db.branch, i.source, i.no_faktur_pgi, i.imei_sn, ik.name as kind_name, ib.name as brand_name, " +
		"it.name as type_name, i.year, i.pawned_at, i.status, i.grade_pgi, " +
		"(CASE WHEN i.is_batangan_pgi = 1 THEN 'Batangan' ELSE 'Lengkap' END) as batangan_pgi, i.price_at_pawn, " +
		"i.base_price, i.final_price, i.capital, i.grade, (CASE WHEN isp.id IS NULL then '-' ELSE isp.name END) as spec_name, " +
		"(CASE WHEN i.is_batangan IS NOT NULL THEN 'Batangan' ELSE 'Lengkap' END) as batangan, " +
		"(CASE WHEN w.id IS NULL THEN '-' ELSE w.name END) as warehouse_name, i.missing_accessories, i.not_ori_accessories, " +
		"(CASE WHEN u.id IS NULL THEN '-' ELSE u.fullname END) as fullname, i.approved_at, i.description " +
		"FROM inventories i " +
		"JOIN (SELECT item_kind_id " +
		"FROM divisions_item_kinds " +
		"WHERE division_id != ? " +
		"AND division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?)) AS dik " +
		"JOIN delivery_batches db ON i.delivery_batch_id =db.id " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN item_types it ON i.type_id = it.id " +
		"JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN item_kinds ik ON i.item_kind_id = ik.id " +
		"LEFT JOIN users u ON i.approved_by = u.id " +
		"LEFT JOIN item_specs isp ON i.item_spec_id = isp.id " +
		"LEFT JOIN warehouses w ON i.last_warehouse_id = w.id " +
		"WHERE i.first_approved_at BETWEEN ? AND ? " +
		"AND i.is_deleted = 0 " +
		"AND db.delivery_id IN (?)"

	QueryExportDeliveryItemReport = "SELECT d.date, db.branch, i.source, i.no_faktur_pgi, i.imei_sn, ik.name as kind_name, ib.name as brand_name, " +
		"it.name as type_name, i.year, i.pawned_at, i.status, i.grade_pgi, " +
		"(CASE WHEN i.is_batangan_pgi = 1 THEN 'Batangan' ELSE 'Lengkap' END) as batangan_pgi, i.price_at_pawn, " +
		"i.base_price, i.final_price, i.capital, i.grade, (CASE WHEN isp.id IS NULL then '-' ELSE isp.name END) as spec_name, " +
		"(CASE WHEN i.is_batangan IS NOT NULL THEN 'Batangan' ELSE 'Lengkap' END) as batangan, " +
		"(CASE WHEN w.id IS NULL THEN '-' ELSE w.name END) as warehouse_name, i.missing_accessories, i.not_ori_accessories, " +
		"(CASE WHEN u.id IS NULL THEN '-' ELSE u.fullname END) as fullname, i.approved_at, i.description " +
		"FROM inventories i " +
		"JOIN (SELECT item_kind_id " +
		"FROM divisions_item_kinds " +
		"WHERE division_id != ? " +
		"AND division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?)) AS dik " +
		"JOIN delivery_batches db ON i.delivery_batch_id =db.id " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN item_types it ON i.type_id = it.id " +
		"JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN item_kinds ik ON i.item_kind_id = ik.id " +
		"LEFT JOIN users u ON i.approved_by = u.id " +
		"LEFT JOIN item_specs isp ON i.item_spec_id = isp.id " +
		"LEFT JOIN warehouses w ON i.last_warehouse_id = w.id " +
		"WHERE i.is_deleted = 0 " +
		"AND db.delivery_id IN (?)"
)
