class Order {
  final int id;
  final String code;
  final String total;
  final String note;
  final String paymentMethod;
  final int tableId;
  final String pickUpType;
  final String proofOfPayment;
  final int status;

  Order({
    required this.id,
    required this.code,
    required this.total,
    required this.note,
    required this.paymentMethod,
    required this.tableId,
    required this.pickUpType,
    required this.proofOfPayment,
    required this.status,
  });

  factory Order.fromJson(Map<String, dynamic> json) {
    return Order(
      id: json['id'] ?? 0,
      code: json['code'] ?? '',
      total: json['total'] ?? '',
      note: json['note'] ?? '',
      paymentMethod: json['payment_method'] ?? '',
      tableId: json['table_id'] ?? 0,
      pickUpType: json['pick_up_type'] ?? '',
      proofOfPayment: json['proof_of_payment'] ?? '',
      status: json['status'] ?? 0,
    );
  }
}
