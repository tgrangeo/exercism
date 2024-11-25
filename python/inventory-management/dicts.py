"""Functions to keep track and alter inventory."""


def create_inventory(items):
    return add_items({},items)

def add_items(inventory, items):
    for i in items:
        inventory[i] = inventory.get(i, 0) + 1
    return inventory


def decrement_items(inventory, items):
    for i in items:
        if inventory.get(i, 0) and inventory.get(i, 0) > 0:
            inventory[i] = inventory.get(i, 0) - 1
    return inventory


def remove_item(inventory, item):
    if inventory.get(item, 0):
        inventory.pop(item)
    return inventory
    

def list_inventory(inventory):
    list = []
    for k,v in inventory.items():
        if v != 0:
            list.append((k,v))
    return list

