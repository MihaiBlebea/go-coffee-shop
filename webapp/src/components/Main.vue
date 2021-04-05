<template>
    <div class="card">
        <div class="card-body">
            <div class="row">
                <div class="col-md-4">
                    <div class="d-flex justify-content-between px-3">
                        <div>Name</div>
                        <div>Size</div>
                        <div>Price</div>
                    </div>
                    <ul class="list-group">
                        <li 
                            class="list-group-item clickable" 
                            v-for="(coffee, index) in menu" 
                            :key="index"
                            v-on:click="selectCoffee(index)"
                        >
                            <div class="d-flex justify-content-between">
                                <div>{{ coffee.name }}</div>
                                <div>{{ sizeFormat(coffee.size) }}</div>
                                <div>{{ moneyFormat(coffee.price) }}</div>
                            </div>
                        </li>
                    </ul>

                    <hr/>

                    <ul class="list-group">
                        <li 
                            class="list-group-item clickable" 
                            v-for="(order, index) in orders" 
                            :key="index"
                        >
                            <div class="d-flex justify-content-between">
                                <div>{{ orderIdFormat(order.id) }}</div>
                                <div>{{ moneyFormat(coffee.price) }}</div>
                            </div>
                        </li>
                    </ul>
                </div>
                <div class="col-md-8">
                    <div v-if="coffee">
                        <h5 class="card-title d-flex justify-content-between">
                            <div>{{ coffee.name }} - {{ sizeFormat(coffee.size) }}</div>
                            <div>{{ moneyFormat(coffee.price) }}</div>     
                        </h5>
                        <img class="w-100 mb-3" :src="coffee.image" />
                        <a v-on:click="order()" class="btn btn-primary">Order {{ coffee.name }}</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    name: 'Main',
    data: function() {
        return {
            accountId: null,
            menu: [],
            coffee: null,
            orders: null
        }
    },
    methods: {
        setupMenu: async function() {
            try {
                let result = await axios.get("http://localhost:8081/api/v1/menu")
                this.menu = result.data.coffees
            } catch(err) {
                console.error(err)
            }
        },
        obtainAccountId: async function() {
            if (! localStorage.getItem("accountId")) {
                try {
                    let result = await axios.post("http://localhost:8083/api/v1/user", {
                        first_name: "fofo",
                        last_name: "bobo",
                        age: 32
                    })

                    this.accountId = result.data.id
                    localStorage.setItem("accountId", result.data.id)
                } catch(err) {
                    console.error(err)
                }
            } else {
                this.accountId = localStorage.getItem("accountId")
            }
        },
        fetchOrders: async function() {
            try {
                let result = await axios.get("http://localhost:8081/api/v1/orders?user_id=" + this.accountId)

                if (result.data.success === true) {
                    this.orders = result.data.orders
                }
            } catch(err) {
                console.error(err)
            }
        },
        selectCoffee: function(index) {
            let coffee = this.menu[index]
            this.coffee = coffee
        },
        moneyFormat: function(money) {
            return "£"  + (money / 100).toFixed(2)
        },
        sizeFormat: function(size) {
            switch(size) {
                case 0:
                    return "S"
                case 1:
                    return "M"
                case 2:
                    return "L"
                default:
                    return "N/A"
            }
        },
        orderIdFormat: function(id) {
            return id.substr(id.length - 5)
        },
        order: async function() {
            try {
                let result = await axios.post("http://localhost:8081/api/v1/order", {
                    coffee_id: this.coffee.id,
                    user_id: this.accountId,
                    amount: this.coffee.price
                })
                console.log(result)
                this.coffee = null
            } catch(err) {
                console.error(err)
            }
        }
    },
    mounted: async function() {
        await this.obtainAccountId()
        await this.setupMenu()
        await this.fetchOrders()
    }
}
</script>

<style scoped>
.card {
    background-color: #3B3B3B;
    color: #FFF;
}

.clickable {
    cursor: pointer;
}
</style>

