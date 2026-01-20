## Documentation reference (Readme.io):

Reference to entities etc: https://docs.ekko.earth/v3/reference

## Folder structure

Structure domains as folders, which contain sub-folders, representing sub-domains, that actually contain the various files (e.g. entity.o, handler.go, repo.go, etc.)

## Software Architecture

Design a “future-microservices-friendly monolith” approach for our flow, without over-engineering, by doing all of the following:
- DDD for boundaries 
- Hexagonal (Ports/Adapters) inside each boundary
- Vertical Slice Architecture for the request flow.

Those three aren’t mutually exclusive—they solve different problems:

- DDD answers: “Where do responsibilities live? Who owns what? How do we prevent tight coupling?”
- Hexagonal answers: “How do we isolate domain logic from DB/APIs so extraction later is easy?”
- Vertical Slice answers: “How do we implement a feature end-to-end without touching 40 files?”

## Feature to implement: Quote creation

In order to create our quote in our system, these are the sequence steps:

1. Get or Validate Organisation (Organisation Domain - organisation sub Domain) 
    1.  if the header org id = to the body organisation Id then move on. If not then
    2. Check if the body organisation Id is a child of the header organisation. if it is then we can move to the next step if not then return a message that this is not allowed.
    3. Return the org details
2. Get or Create a new Customer (Organisation Domain - Customer sub Domain)
3. Calculate carbon foot print (Impact Domain - Carbon Footprint sub Domain)
    1. If not merchant details is included in the body of the request use the mcc details from organisation
    2. We need to get the merchant country id that is equal to the country code passed in the request body. (Platform Domain - country sub Domain)
    3. if the currency is not euro we need to convert the currency by getting the currency conversion rate from the currency conversion rate table. (finance domain - currency subdomain) *** needs to conclude where this is owned
    4. with the country id and mcc code and  we will look up the factor from the carbon factor table.
    5. Then calculate the carbon footprint using the transactions amount from the body
    6. Write to footprint
    7. return the footprint 
4. Get blended project unit price (Impact Partner Domain)
    1. Get all impact partners for organisation (impact partner sub Domain)
    2. Get projects for each impact partners and filter by location if true(Impact partner domain - projects sub domain)
    3. Create blended price projects ( impact partner domain - projects sub domain)
    4. Return object
5. Calculate Compensation Amount (Impact domain - quote sub domain)
6. Calc Round Up (Impact domain Quote Sub Domain)
7. Calculate Service Fee (Impact domain - Fee Sub Domain) 
8. Calc Sales Tax (Funds domain - sales tax sub domain)
    1. mcc country, state , postal and customer country state postal code is needed to get it from either the sales tax table or a api call to sales tax provider to be the factor
    2. need to get data from country/ platform domain
    3. calc and return sales tax factor

9.  Write to Quote

Questions to Ask

- Can any steps run in parallel?
- What happens if an intermediate step fails?
- How do you handle partial rollback?
- Who owns the schema migrations?
- Can you deploy domain A without domain B?
- How do you handle schema changes that affect multiple domains?
- How many files does a developer need to touch?
- How easy is it for a new developer to understand the flow?
- Can you trace a request through the system?
- If you change the footprint calculation, how many places break?
- Who owns each domain?
- How do teams coordinate changes?
- When we extract to micro services, are we comfortable with eventual consistency (data replication) or do we need strong consistency (RPC)?