class Order {
  final int id;
  final int code;
  final int total;
  final String note;
  final String PaymentMethod;
  final int tabelId;
  final String PickUptype;
  final String ProofOfPayment;
  final int status;

  Order({
    required this.id,
    required this.code,
    required this.total,
    required this.note,
    required this.PaymentMethod,
    required this.tabelId,
    required this.PickUptype,
    required this.ProofOfPayment,
    required this.status,
  });

  factory Order.fromJson(Map<String, dynamic> json) {
    return Order(
      id: json['id'],
      code: json['code'],
      total: json['total'],
      note: json['note'],
      PaymentMethod: json['payment_method'],
      tabelId: json['table_id'],
      PickUptype: json['pick_up_type'],
      ProofOfPayment: json['proof_of_payment'],
      status: json['status'],
    );
  }
}
