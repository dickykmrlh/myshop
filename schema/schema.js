const graphql = require('graphql');
const axios = require('axios');
const {
    GraphQLObjectType,
    GraphQLString,
    GraphQLInt,
    GraphQLFloat,
    GraphQLSchema
} = graphql;

const ruleType = new GraphQLObjectType({
    name: 'rule',
    fields: {
        minimum_qty: { type: GraphQLInt },
        must_buy: { type: GraphQLString },
    }
});

const PromotionType = new GraphQLObjectType({
    name: 'Promotion',
    fields: {
        sku: { type: GraphQLString },
        type: { type: GraphQLString },
        age: { type: GraphQLInt },
        rule: {
            type: ruleType,
            resolve(parentValue, args) {
                return axios.get(`http://localhost:3000/companies/${parentValue.sku}`)
                    .then(res => res.data);
            }
        }
    }
});

const InventoryType = new GraphQLObjectType({
    name: 'Inventory',
    fields: {
        sku: { type: GraphQLString },
        name: { type: GraphQLString },
        price: { type: GraphQLFloat },
        qty: {type: GraphQLInt},
        promotion: {
            type: PromotionType,
            resolve(parentValue, args) {
                return axios.get(`http://localhost:3000/companies/${parentValue.sku}`)
                    .then(res => res.data);
            }
        }
    }
});

const RootQuery = new GraphQLObjectType({
    name: 'RootQueryType',
    fields: {
        inventory: {
            type: InventoryType,
            args: { id: { type: GraphQLString } },
            resolve(parentValue, args) {
                return axios.get(`http://localhost:3000/users/${args.id}`)
                    .then(resp => resp.data);
            }
        },
        promotion: {
            type: PromotionType,
            args: { id: {type: GraphQLString } },
            resolve(parentValue, args) {
                return axios.get(`http://localhost:3000/companies/${args.id}`)
                    .then(resp => resp.data);
            }
        }
    }
});

module.exports = new GraphQLSchema({
    query: RootQuery
});