# there are three important variable
#
# boxes = [3, 1, 6]-----------number of boxes
# unit_per_box = [2, 7, 4]-----------number of units per box
# there are 3 box with 2 unit ,1 box with 7 unit ,6 box with 4 unit
# truck_size = 6----------number of box truck can carry
# truck can carry 6 box

# we have to maximise no.of unit and not the number of box
# so we will keep 1 box with 7 unit, and 5 box with 4 unit,
# a total of 6 box with 27 unit,
# to do this we will
#   -first find the box with highest number of unit per box
#   -then keep those boxes in the truck and
#   -repeat until the truck is full
#

# boxes = [3, 1, 6]
# unit_per_box = [2, 7, 4]
# truck_size = 6
num = int(input())
boxes = [int(i) for i in input().replace('[', "").replace(']', '').split(',')]
unitSize=int(input())
unit_per_box = [int(i) for i in input().replace('[', "").replace(']', '').split(',')]
truck_size = int(input())

units_in_truck = 0
available_space_in_truck = truck_size

for number_of_elements in range(len(boxes)):  # this is the repeat part
    pos_of_box_with_most_unit = 0  # this is initially 0
    for i in range(len(unit_per_box)):          # liniar search to find box with most unit
        if unit_per_box[i] > unit_per_box[pos_of_box_with_most_unit]:
            pos_of_box_with_most_unit = i

    if available_space_in_truck == 0:  # check if truck is already full
        break
    else:
        # check if all the box can fit in
        if available_space_in_truck >= boxes[pos_of_box_with_most_unit]:
            #update number of units in truck
            units_in_truck = units_in_truck + \
                (boxes[pos_of_box_with_most_unit] *
                 unit_per_box[pos_of_box_with_most_unit])
            #update the amount of space left
            available_space_in_truck = available_space_in_truck - \
                boxes[pos_of_box_with_most_unit]
            #make the element we filled zero since we dont want to find it again
            unit_per_box[pos_of_box_with_most_unit] = 0
        else:  # only fill in the available_space_in_truck
            units_in_truck = units_in_truck + \
                (available_space_in_truck *
                 unit_per_box[pos_of_box_with_most_unit])
            #now truck is full
            available_space_in_truck = 0
            unit_per_box[pos_of_box_with_most_unit] = 0

#finally print the answer
print(units_in_truck)
