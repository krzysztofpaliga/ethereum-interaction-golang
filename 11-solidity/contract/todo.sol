// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <=0.9.0;

contract ToDo {
    address public owner;

    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isOwner() {
        require(owner == msg.sender);
        _;
    }
    function add(string memory _content) public isOwner {
        tasks.push(Task(_content, false));
    }

    function get(uint _id) public view returns (Task memory) {
        return tasks[_id];
    }

    function list() public view returns (Task[] memory) {
        return tasks;
    }

    function update(uint _id, string memory _content) public isOwner {
        tasks[_id].content = _content;
    }

    function toggle(uint _id) public isOwner {
        tasks[_id].status = !tasks[_id].status;
    }

    function remove(uint _id) public isOwner {
        require(owner == msg.sender);
        for (uint i = _id; i < tasks.length - 1; i++) {
            tasks[i] = tasks[i + 1];
        }
        tasks.pop();
    }
}
