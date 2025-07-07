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

	QueryExportDeliveryBatchGoldEXTReport = "SELECT d.id, d.date, db.branch, d.source, inv.item_count " +
		"FROM delivery_batches db " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN (SELECT i.delivery_batch_id, COUNT(*) as item_count " +
		"FROM inventories i " +
		"JOIN divisions_item_kinds dik ON i.item_kind_id = dik.item_kind_id " +
		"WHERE dik.division_id = ? " +
		"AND dik.division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?) " +
		"GROUP BY i.delivery_batch_id) AS inv ON inv.delivery_batch_id = db.id " +
		"WHERE d.source = 'EXT' " +
		"ORDER BY BRANCH"

	QueryExportDeliveryBatchGoldReport = "SELECT d.id, d.date, db.branch, d.source, inv.item_count " +
		"FROM delivery_batches db " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN (SELECT i.delivery_batch_id, COUNT(*) as item_count " +
		"FROM inventories i " +
		"JOIN divisions_item_kinds dik ON i.item_kind_id = dik.item_kind_id " +
		"WHERE dik.division_id = ? " +
		"AND dik.division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?) " +
		"GROUP BY i.delivery_batch_id) AS inv ON inv.delivery_batch_id = db.id " +
		"WHERE d.source != 'EXT' " +
		"AND d.date BETWEEN ? AND ? " +
		"ORDER BY BRANCH"

	QueryExportDeliveryItemGoldEXTReport = "SELECT d.date, db.branch, i.source, i.no_faktur_pgi, i.imei_sn, ik.name as kind_name, ib.name as brand_name, it.name as type_name, " +
		"ig.purity, ig.dry_weight, ig.weight_reduction, " +
		"(CASE WHEN ig.weight_reduction IS NULL THEN '-' ELSE CAST(ig.dry_weight - ig.weight_reduction AS CHAR) END) AS net_weight, " +
		"gmm.name as gold_mint_mark_name, gt.name, ig.piece_count, i.status, " +
		"(CASE WHEN w.id IS NULL THEN '-' ELSE w.name END) as warehouse_name, " +
		"(CASE WHEN gmm.id IS NOT NULL THEN 'Ya' ELSE 'Tidak' END ) as is_cap, " +
		"(CASE WHEN u.id IS NULL THEN '-' ELSE u.fullname END) as fullname, " +
		"i.approved_at, i.description, i.item_kind_id, i.type_id " +
		"FROM inventories i " +
		"JOIN (SELECT item_kind_id " +
		"FROM divisions_item_kinds " +
		"WHERE division_id = ? " +
		"AND division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?)) AS dik ON dik.item_kind_id = i.item_kind_id " +
		"JOIN delivery_batches db ON i.delivery_batch_id = db.id " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN item_types it ON i.type_id = it.id " +
		"JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN item_kinds ik ON i.item_kind_id = ik.id " +
		"LEFT JOIN users u ON i.approved_by = u.id " +
		"LEFT JOIN item_specs isp ON i.item_spec_id = isp.id " +
		"LEFT JOIN warehouses w ON i.last_warehouse_id = w.id " +
		"LEFT JOIN inventory_golds ig ON ig.inventory_id = i.id " +
		"LEFT JOIN gold_mint_marks gmm ON ig.gold_mint_mark_id = gmm.id " +
		"LEFT JOIN gold_types gt ON ig.gold_type_id = gt.id " +
		"WHERE i.first_approved_at BETWEEN ? AND ? " +
		"AND i.is_deleted = 0 " +
		"AND db.delivery_id IN (?)"

	QueryExportDeliveryItemGoldReport = "SELECT d.date, db.branch, i.source, i.no_faktur_pgi, i.imei_sn, ik.name as kind_name, ib.name as brand_name, it.name as type_name, " +
		"ig.pgi_purity, ig.pgi_dry_weight, ig.pgi_weight_reduction, " +
		"(CASE WHEN ig.weight_reduction IS NULL THEN '-' ELSE CAST(ig.dry_weight - ig.weight_reduction AS CHAR) END) AS net_weight, " +
		"gmm.name as gold_mint_mark_name, gt.name, ig.piece_count, i.pawned_at, i.status, " +
		"i.grade_pgi, i.price_at_pawn, " +
		"(CASE WHEN w.id IS NULL THEN '-' ELSE w.name END) as warehouse_name, " +
		"(CASE WHEN gmm.id IS NOT NULL THEN 'Ya' ELSE 'Tidak' END ) as is_cap, " +
		"(CASE WHEN u.id IS NULL THEN '-' ELSE u.fullname END) as fullname, " +
		"i.approved_at, i.description, i.item_kind_id, i.type_id " +
		"FROM inventories i " +
		"JOIN (SELECT item_kind_id " +
		"FROM divisions_item_kinds " +
		"WHERE division_id = ? " +
		"AND division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?)) AS dik ON dik.item_kind_id = i.item_kind_id " +
		"JOIN delivery_batches db ON i.delivery_batch_id = db.id " +
		"JOIN deliveries d ON db.delivery_id = d.id " +
		"JOIN item_types it ON i.type_id = it.id " +
		"JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN item_kinds ik ON i.item_kind_id = ik.id " +
		"LEFT JOIN users u ON i.approved_by = u.id " +
		"LEFT JOIN item_specs isp ON i.item_spec_id = isp.id " +
		"LEFT JOIN warehouses w ON i.last_warehouse_id = w.id " +
		"LEFT JOIN inventory_golds ig ON ig.inventory_id = i.id " +
		"LEFT JOIN gold_mint_marks gmm ON ig.gold_mint_mark_id = gmm.id " +
		"LEFT JOIN gold_types gt ON ig.gold_type_id = gt.id " +
		"AND i.is_deleted = 0 " +
		"AND db.delivery_id IN (?)"

	QueryExportInventoryMovementReportExport = "SELECT im.id, im.no_ref, im.no_faktur, im.move_date, im.confirmed_at, u.fullname, " +
		"(CASE WHEN im.status = 2 THEN 'CONFIRMED' ELSE 'DRAFT' END) as status, " +
		"wf.name as from_name, owf.name as from_office_name, wt.name as to_name, owt.name as to_office_name, " +
		"imd.details_count, im.note " +
		"FROM inventory_movements im " +
		"JOIN users u ON u.id = im.confirmed_by " +
		"JOIN (SELECT imd.inventory_movement_id, COUNT(*) as details_count from inventory_movement_details imd GROUP BY imd.inventory_movement_id) imd ON im.id = imd.inventory_movement_id " +
		"LEFT JOIN warehouses wf ON wf.id = im.from " +
		"LEFT JOIN warehouses wt ON wt.id = im.to " +
		"LEFT JOIN offices owf ON owf.id = wf.office_id " +
		"LEFT JOIN offices owt ON owt.id = wt.office_id " +
		"WHERE status = 2 AND division_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?) " +
		"AND move_date BETWEEN ? AND ? " +
		"AND wf.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"AND wt.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) "

	QueryInventoryMovementItemReportExport = "SELECT imd.id, im.no_faktur, im.no_ref, i.no_faktur_pgi, i.imei_sn, " +
		"(CASE WHEN i.pawned_at IS NULL THEN '-' ELSE i.pawned_at END) as pawned_at, " +
		"i.created_at, wf.name as from_name, owf.name as office_from_name, wt.name as to_name, " +
		"owt.name as office_to_name, imd.kind, imd.brand, imd.type, i.year, " +
		"(CASE WHEN isp.name IS NULL THEN '-' ELSE isp.name END) as spec_name, " +
		"(CASE WHEN i.is_batangan IS NULL THEN 'Lengkap' ELSE 'Batangan' END) as batangan, " +
		"i.grade, i.grade_pgi, i.price_at_pawn, i.capital " +
		"FROM inventory_movement_details imd " +
		"JOIN inventory_movements im ON imd.inventory_movement_id = im.id " +
		"JOIN inventories i ON imd.inventory_id = i.id " +
		"JOIN item_specs isp ON i.item_spec_id = isp.id " +
		"LEFT JOIN warehouses wf ON wf.id = im.from " +
		"LEFT JOIN warehouses wt ON wt.id = im.to " +
		"LEFT JOIN offices owf ON owf.id = wf.office_id " +
		"LEFT JOIN offices owt ON owt.id = wt.office_id " +
		"WHERE imd.inventory_movement_id IN (?) "

	QueryCatalogCustomerLoginLogsReport = "SELECT cc.name, cc.shop_name, cc.email, cc.handphone, cc.address, ccll.login_at " +
		"FROM catalog_customer_login_logs ccll " +
		"JOIN catalog_customers cc ON ccll.catalog_customer_id = cc.id " +
		"WHERE ccll.login_at BETWEEN CONCAT(?, ' 00:00:00') AND CONCAT(?, ' 23:59:59')"

	QueryExportUserTaskCountLogReport = "SELECT utcl.id, utcl.log_date, u.fullname, utcl.approve_count, " +
		"utcl.invoice_detail_count, utcl.return_count, utcl.request_reset_count, " +
		"utcl.adjustment_count, utcl.cetak_barcode, utcl.pindah_gudang, utcl.upload_foto, utcl.input_aksesoris " +
		"FROM user_task_count_logs utcl " +
		"JOIN users_divisions ud ON ud.user_id IN (SELECT ud.division_id FROM users_divisions ud WHERE user_id = ?) " +
		"JOIN users u ON utcl.user_id = u.id " +
		"WHERE utcl.log_date BETWEEN CONCAT(?, ' 00:00:00') AND CONCAT(?, ' 23:59:59') "

	QueryExportInventoryReturnsReport = "SELECT inv.no_faktur_pgi, ir.created_at, ir.request_note, ir.processed_at, u.fullname as processed_name, " +
		"ir.cancel_at, ir.cancel_note, u2.fullname as cancel_name, " +
		"(CASE WHEN ir.status = 0 THEN 'Pending' WHEN ir.status = 1 THEN 'Pengiriman' WHEN ir.status = 2 THEN 'Tiba di PGI' " +
		"WHEN ir.status = 9 THEN 'Ditolak' WHEN ir.status = 10 THEN 'Batal' ELSE '-' END) as status " +
		"FROM inventory_returns ir " +
		"LEFT JOIN inventories inv ON ir.inventory_id = inv.id " +
		"LEFT JOIN users u ON ir.processed_by = u.id " +
		"LEFT JOIN users u2 ON ir.cancelled_by = u.id " +
		"WHERE ir.created_at BETWEEN CONCAT(?, ' 00:00:00') AND CONCAT(?, ' 23:59:59') " +
		"AND inv.office_id IN (SELECT office_id FROM users_offices WHERE user_id = ?) " +
		"ORDER BY ir.id"

	QueryExportSendbackReport = "SELECT s.id, s.no_faktur, s.send_date, s.confirmed_at, u.fullname, " +
		"(CASE WHEN s.status = 2 THEN 'CONFIRMED' ELSE 'DRAFT' END) as status, " +
		"w.name, d.details_count, s.notes " +
		"FROM sendbacks s " +
		"JOIN warehouses w ON s.warehouse_id = w.id " +
		"JOIN users u ON s.confirmed_by = u.id " +
		"JOIN (SELECT sendback_id, COUNT(*) as details_count FROM sendback_details GROUP BY sendback_id) as d ON d.sendback_id = s.id " +
		"WHERE s.status = 2 " +
		"AND s.division_id IN (SELECT division_id FROM users_divisions WHERE user_id = ?) " +
		"AND s.send_date BETWEEN ? AND ? " +
		"ORDER BY confirmed_at"

	QueryExportSendbackDetailReport = "SELECT s.no_faktur, inv.no_faktur_pgi, inv.created_at, ik.name as kind_name, inv.imei_sn, " +
		"ib.name as brand_name, it.name as type_name, COALESCE(isp.name, '-') as spec_name, " +
		"inv.year, (CASE WHEN inv.is_batangan IS NOT NULL THEN 'Batangan' ELSE 'Lengkap' END) as batangan, " +
		"inv.grade, inv.grade_pgi, inv.final_price_after_adj, w.name as warehouse_name, adj.adj_name " +
		"FROM sendback_details sd " +
		"JOIN sendbacks s ON sd.sendback_id = s.id " +
		"JOIN inventories inv ON sd.inventory_id = inv.id " +
		"JOIN item_kinds ik ON inv.item_kind_id = ik.id " +
		"JOIN item_types it ON inv.type_id = it.id " +
		"JOIN item_brands ib ON it.item_brand_id = ib.id " +
		"JOIN item_specs isp ON inv.item_spec_id = isp.id " +
		"JOIN warehouses w ON w.id = s.warehouse_id " +
		"LEFT JOIN (SELECT dtl.id, GROUP_CONCAT(adj.name SEPARATOR ', ') as adj_name " +
		"FROM inventories dtl " +
		"JOIN inventory_invoice_adjustments adj ON dtl.id = adj.inventory_id " +
		"GROUP BY dtl.id) adj ON adj.id = inv.id " +
		"WHERE s.id IN (?)"
)
