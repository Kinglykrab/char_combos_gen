# Character Creation Combination Generator
This is a Go program I wrote just for fun to convert my old Perl script to Go.

To use this you will need [Go](https://go.dev/) installed and will have to run `go run combinations.go` from a CLI of your choice.

## Constants
These constants are used in the `char_create_combinations` and `start_zones` queries and will need to be changed.

| Constant | Description |
| :--- | :--- |
| startZoneID | Zone ID |
| startX | X Coordinate |
| startY | Y Coordinate |
| startZ | Z Coordinate |
| startHeading | Heading Coordinate |

## Query Data
A singular allocation with an ID of `0` with all base stats at `100` is created for `char_create_point_allocations`.

Truncate queries are at the top of the generated file to delete pre-existing data for insertion.

Sets of queries are generated with race and class descriptors at the end of the query for easy reading.

## Query Images
![image](https://user-images.githubusercontent.com/89047260/172984382-2dbaf597-b2e6-4f1e-bdb3-55cf20049892.png)
![image](https://user-images.githubusercontent.com/89047260/172984429-b597de90-a831-4c6d-b370-1f9c2c95cbcc.png)

## Notes
This generator sets all the character creation combinations to the same zone ID and starting coordinates.

This generator as is does not allow for non-caster Races to be caster classes since these classes do not have Robe models, meaning an Ogre cannot be a Magician for example.

This generator as is does not allow for non-Monk Races to be Monk, meaning an Erudite cannot be a Monk.
