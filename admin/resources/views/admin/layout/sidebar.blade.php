<ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" id="accordionSidebar">

    <!-- Sidebar - Brand -->
    <a class="sidebar-brand d-flex align-items-center justify-content-center" href="dashboard">
        <div class="sidebar-brand-icon rotate">
            <i class="fas fa-user"></i>
        </div>
        <div class="sidebar-brand-text mx-3">{{ $data['message']['username'] }}</div>
    </a>

    <!-- Divider -->
    <hr class="sidebar-divider my-0">

    <!-- Nav Item - Dashboard -->
    <li class="nav-item">
        <a class="nav-link" href="/admin/dashboard">
            <i class="fas fa-fw fa-tachometer-alt"></i>
            <span>Dashboard</span></a>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider">

    <!-- Heading -->
    <div class="sidebar-heading">
        Manage
    </div>

    <!-- Nav Item - Pages Collapse Menu -->
    <li class="nav-item">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#category" aria-expanded="true"
            aria-controls="category">
            <i class="fas fa-fw fa-list"></i>
            <span>Category</span>
        </a>
        <div id="category" class="collapse" aria-labelledby="headingTwo" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Category</h6>
                <a class="collapse-item" href="/admin/category/create">Create Category</a>
                <a class="collapse-item" href="/admin/category">List Category</a>
            </div>
        </div>
    </li>

    <!-- Nav Item - Utilities Collapse Menu -->
    <li class="nav-item">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#Product" aria-expanded="true"
            aria-controls="Product">
            <i class="fas fa-fw fa-book"></i>
            <span>Product</span>
        </a>
        <div id="Product" class="collapse" aria-labelledby="headingUtilities" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Product</h6>
                <a class="collapse-item" href="/admin/product/create">Create Product</a>
                <a class="collapse-item" href="/admin/product">List Product</a>
            </div>
        </div>
    </li>


    <!-- Nav Item - Utilities Collapse Menu -->
    <li class="nav-item">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#Table" aria-expanded="true"
            aria-controls="Product">
            <i class="fas fa-fw fa-book"></i>
            <span>Table</span>
        </a>
        <div id="Table" class="collapse" aria-labelledby="headingUtilities" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Table</h6>
                <a class="collapse-item" href="/admin/table/create">Create Table</a>
                <a class="collapse-item" href="/admin/table">List Table</a>
            </div>
        </div>
    </li>

    <!-- Nav Item - Utilities Collapse Menu -->
    <li class="nav-item">
        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#Order" aria-expanded="true"
            aria-controls="Product">
            <i class="fas fa-fw fa-book"></i>
            <span>Order</span>
        </a>
        <div id="Order" class="collapse" aria-labelledby="headingUtilities" data-parent="#accordionSidebar">
            <div class="bg-white py-2 collapse-inner rounded">
                <h6 class="collapse-header">Custom Order</h6>
                {{-- <a class="collapse-item" href="/admin/Order/create">Create Order</a> --}}
                <a class="collapse-item" href="/admin/order">List Order</a>
            </div>
        </div>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider d-none d-md-block">

    <!-- Sidebar Toggler (Sidebar) -->
    <div class="text-center d-none d-md-inline">
        <button class="rounded-circle border-0" id="sidebarToggle"></button>
    </div>


</ul>