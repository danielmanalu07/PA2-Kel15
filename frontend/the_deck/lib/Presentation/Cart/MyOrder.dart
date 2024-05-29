import 'dart:io';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:image_picker/image_picker.dart';
import 'package:the_deck/Controller/CustomerController.dart';
import 'package:the_deck/Models/Order.dart';

class MyOrder extends StatefulWidget {
  const MyOrder({Key? key}) : super(key: key);

  @override
  State<MyOrder> createState() => _MyOrderState();
}

class _MyOrderState extends State<MyOrder> {
  List<Order> orders = [];
  bool isLoading = true;
  final RegisterController _customerController = Get.find<RegisterController>();
  final ImagePicker _picker = ImagePicker();
  XFile? _image;

  @override
  void initState() {
    super.initState();
    _fetchOrders();
  }

  Future<void> _fetchOrders() async {
    try {
      await _customerController.getMyOrder();
      setState(() {
        orders = _customerController.orderItems;
        isLoading = false;
      });
    } catch (e) {
      print('Error: $e');
      setState(() {
        isLoading = false;
      });
    }
  }

  Future<void> _pickImage(BuildContext context) async {
    final pickedFile = await _picker.pickImage(source: ImageSource.gallery);
    if (pickedFile != null) {
      setState(() {
        _image = pickedFile;
      });
    }
  }

  void _showImagePickerDialog(BuildContext context, int orderId) {
    showDialog(
      context: context,
      builder: (context) => Dialog(
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(12.0),
        ),
        child: Container(
          padding: EdgeInsets.all(16.0),
          child: SingleChildScrollView(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Upload Proof of Payment',
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 20,
                  ),
                ),
                SizedBox(height: 16),
                Container(
                  width: double.infinity,
                  height: 450,
                  decoration: BoxDecoration(
                    color: Colors.grey[200],
                    borderRadius: BorderRadius.circular(8),
                    border: Border.all(color: Colors.grey[400]!),
                  ),
                  child: Image.asset(
                    'assets/img/QRIS.jpeg', // Path to your QRIS image
                    fit: BoxFit.cover,
                  ),
                ),
                SizedBox(height: 16),
                GestureDetector(
                  onTap: () async {
                    await _pickImage(context);
                    setState(() {});
                  },
                  child: Container(
                    width: double.infinity,
                    height: 150,
                    decoration: BoxDecoration(
                      color: Colors.grey[200],
                      borderRadius: BorderRadius.circular(8),
                      border: Border.all(color: Colors.grey[400]!),
                    ),
                    child: _image != null
                        ? Image.file(
                            File(_image!.path),
                            fit: BoxFit.cover,
                          )
                        : Center(
                            child: Text(
                              'Tap to select image',
                              style: TextStyle(color: Colors.grey),
                            ),
                          ),
                  ),
                ),
                SizedBox(height: 16),
                if (_image != null) Text('Selected: ${_image!.name}'),
                SizedBox(height: 16),
                Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: [
                    TextButton(
                      onPressed: () {
                        setState(() {
                          _image = null;
                        });
                        Navigator.of(context).pop();
                      },
                      child: Text('Cancel'),
                    ),
                    TextButton(
                      onPressed: () async {
                        await _customerController.uploadImage(
                            context, orderId, _image);
                        await _fetchOrders();
                        Navigator.of(context).pop();
                      },
                      child: Text('Submit'),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          "My Orders",
          style: TextStyle(
            color: Colors.black,
            fontWeight: FontWeight.bold,
          ),
        ),
        backgroundColor: Colors.white,
        elevation: 0,
      ),
      body: isLoading
          ? Center(child: CircularProgressIndicator())
          : orders.isEmpty
              ? Center(child: Text('No orders found.'))
              : ListView.builder(
                  itemCount: orders.length,
                  itemBuilder: (context, index) {
                    final order = orders[index];
                    String status;
                    String table;
                    String payment;
                    if (order.status == 0) {
                      status = 'Waiting';
                    } else if (order.status == 1) {
                      status = 'Accepted';
                    } else if (order.status == 2) {
                      status = 'Rejected';
                    } else if (order.status == 3) {
                      status = 'Ready for Pick Up';
                    } else if (order.status == 4) {
                      status = 'Finished';
                    } else {
                      status = 'Canceled';
                    }
                    if (order.tableId == 0) {
                      table = 'Tidak Menggunakan Meja Makan';
                    } else {
                      table = 'Menggunakan Meja ${order.tableId}';
                    }
                    if (order.proofOfPayment == "") {
                      payment = 'Belum Terkirim';
                    } else {
                      payment = 'Sudah Terkirim';
                    }
                    return Padding(
                      padding: const EdgeInsets.all(8.0),
                      child: Card(
                        elevation: 4,
                        child: ListTile(
                          title: Text('Order Code: ${order.code}'),
                          subtitle: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Text('Total: ${order.total}'),
                              Text('Status: $status'),
                              Text('Jenis Pembayaran: ${order.paymentMethod}'),
                              Text('Table: $table'),
                              Text('Note: ${order.note}'),
                              Text('Jenis Pengambilan: ${order.pickUpType}'),
                              Text('Bukti Pembayaran: $payment'),
                            ],
                          ),
                          trailing: Row(
                            mainAxisSize: MainAxisSize.min,
                            children: [
                              if (order.proofOfPayment.isEmpty &&
                                  status != 'Canceled' &&
                                  order.paymentMethod != "Cash")
                                TextButton(
                                  onPressed: () =>
                                      _showImagePickerDialog(context, order.id),
                                  child: Text(
                                    "Upload Payment",
                                    style: TextStyle(color: Colors.blue),
                                  ),
                                ),
                              if (order.status == 0 &&
                                  order.proofOfPayment.isEmpty)
                                SizedBox(width: 8),
                              if (order.status == 0 &&
                                  order.proofOfPayment.isEmpty)
                                TextButton(
                                  onPressed: () {
                                    _customerController
                                        .updateOrderStatus(order.id);
                                    _fetchOrders();
                                  },
                                  child: Text(
                                    "Cancel Order",
                                    style: TextStyle(
                                      color: Colors.red,
                                      fontSize: 12,
                                    ),
                                  ),
                                ),
                            ],
                          ),
                        ),
                      ),
                    );
                  },
                ),
    );
  }
}
