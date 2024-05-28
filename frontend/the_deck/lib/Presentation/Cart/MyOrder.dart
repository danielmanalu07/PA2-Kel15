import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

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
      builder: (context) => StatefulBuilder(
        builder: (context, setState) {
          return AlertDialog(
            title: Text('Pick Image'),
            content: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
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
                        : Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Icon(
                                Icons.camera_alt,
                                color: Colors.grey[600],
                                size: 50,
                              ),
                              SizedBox(height: 8),
                              Text(
                                "Upload Proof Of Payment Picture",
                                style: TextStyle(
                                  color: Colors.grey[600],
                                  fontSize: 16,
                                ),
                              ),
                            ],
                          ),
                  ),
                ),
                SizedBox(height: 16),
                _image != null
                    ? Text('Selected: ${_image!.name}')
                    : Text('No image selected'),
              ],
            ),
            actions: [
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
                  Navigator.of(context).pop();

                  _fetchOrders();
                },
                child: Text('Submit'),
              ),
            ],
          );
        },
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          "My Order",
          style: TextStyle(
            color: Colors.black,
            wordSpacing: 8,
            fontWeight: FontWeight.bold,
          ),
        ),
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
                    String note;
                    if (order.status == 0) {
                      status = 'Waiting';
                    } else if (order.status == 1) {
                      status = 'Accepted';
                    } else if (order.status == 2) {
                      status = 'Rejected';
                    } else {
                      status = 'Canceled';
                    }
                    if (order.tableId == 0) {
                      table = 'Tidak Menggunakan Meja Makan';
                    } else {
                      table = order.tableId.toString();
                    }
                    if (order.proofOfPayment == "") {
                      payment = 'Belum Terkirim';
                    } else {
                      payment = 'Sudah Terkirim';
                    }
                    if (order.note == "") {
                      note = 'Tidak ada catatan';
                    } else {
                      note = order.note;
                    }
                    return Card(
                      child: ListTile(
                        title: Text('Order Code: ${order.code}'),
                        subtitle: Text(
                            'Total: ${order.total}\nNote: ${note}\nStatus: ${status}\nJenis Pengambilan: ${order.pickUpType}\nJenis Pembayaran: ${order.paymentMethod}\nNo Meja: ${table}\nBukti Pembayaran: ${payment}'),
                        trailing: Row(
                          mainAxisSize: MainAxisSize.min,
                          children: [
                            TextButton(
                              onPressed: () =>
                                  _showImagePickerDialog(context, order.id),
                              child: Text(
                                "Bukti Pembayaran",
                                style: TextStyle(
                                  color: Colors.blue,
                                ),
                              ),
                            ),
                            if (order.status == 0 &&
                                payment == 'Belum Terkirim')
                              TextButton(
                                onPressed: () {
                                  _customerController
                                      .updateOrderStatus(order.id);
                                  _fetchOrders();
                                },
                                child: Text(
                                  "Cancel",
                                  style: TextStyle(
                                    color: Colors.red,
                                  ),
                                ),
                              ),
                          ],
                        ),
                      ),
                    );
                  },
                ),
    );
  }
}
